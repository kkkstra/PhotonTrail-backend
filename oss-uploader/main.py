from flask import Flask, render_template, jsonify
from alibabacloud_tea_openapi.models import Config
from alibabacloud_sts20150401.client import Client as Sts20150401Client
from alibabacloud_sts20150401 import models as sts_20150401_models
from alibabacloud_credentials.client import Client as CredentialClient
import os
import json
import base64
import hmac
import datetime
import time
import hashlib

app = Flask(__name__)

# 配置环境变量 OSS_ACCESS_KEY_ID, OSS_ACCESS_KEY_ID, OSS_STS_ROLE_ARN。
access_key_id = os.environ.get('OSS_ACCESS_KEY_ID')
access_key_secret = os.environ.get('OSS_ACCESS_KEY_SECRET')
role_arn_for_oss_upload = os.environ.get('OSS_STS_ROLE_ARN')

# 自定义会话名称
role_session_name = 'role_session_name'  

# 替换为实际的bucket名称和region_id
bucket = 'photon-trail'
region_id = 'cn-hangzhou'

host = f'https://{bucket}.oss-cn-hangzhou.aliyuncs.com'

# 指定过期时间，单位为秒
expire_time = 3600  
# 指定上传到OSS的文件前缀。
upload_dir = 'img'

def hmacsha256(key, data):
    """
    计算HMAC-SHA256哈希值的函数

    :param key: 用于计算哈希的密钥，字节类型
    :param data: 要进行哈希计算的数据，字符串类型
    :return: 计算得到的HMAC-SHA256哈希值，字节类型
    """
    try:
        mac = hmac.new(key, data.encode(), hashlib.sha256)
        hmacBytes = mac.digest()
        return hmacBytes
    except Exception as e:
        raise RuntimeError(f"Failed to calculate HMAC-SHA256 due to {e}")

@app.route("/")
def hello_world():
    return render_template('index.html')

@app.route('/get_post_signature_for_oss_upload', methods=['GET'])
def generate_upload_params():
    # 初始化配置，直接传递凭据
    config = Config(
        region_id=region_id,
        access_key_id=access_key_id,
        access_key_secret=access_key_secret
    )

    # 创建 STS 客户端并获取临时凭证
    sts_client = Sts20150401Client(config=config)
    assume_role_request = sts_20150401_models.AssumeRoleRequest(
        role_arn=role_arn_for_oss_upload,
        role_session_name=role_session_name
    )
    response = sts_client.assume_role(assume_role_request)
    token_data = response.body.credentials.to_map()

    # 使用 STS 返回的临时凭据
    temp_access_key_id = token_data['AccessKeyId']
    temp_access_key_secret = token_data['AccessKeySecret']
    security_token = token_data['SecurityToken']


    now = int(time.time())
    # 将时间戳转换为datetime对象
    dt_obj = datetime.datetime.utcfromtimestamp(now)
    # 在当前时间增加3小时，设置为请求的过期时间
    dt_obj_plus_3h = dt_obj + datetime.timedelta(hours=3)

    # 请求时间
    dt_obj_1 = dt_obj.strftime('%Y%m%dT%H%M%S') + 'Z'
    # 请求日期
    dt_obj_2 = dt_obj.strftime('%Y%m%d')
    # 请求过期时间
    expiration_time = dt_obj_plus_3h.strftime('%Y-%m-%dT%H:%M:%S.000Z')
    
    # 构建 Policy 并生成签名
    policy = {
        "expiration": expiration_time,
        "conditions": [
            ["eq", "$success_action_status", "200"],
            {"x-oss-signature-version": "OSS4-HMAC-SHA256"},
            {"x-oss-credential": f"{temp_access_key_id}/{dt_obj_2}/cn-hangzhou/oss/aliyun_v4_request"},
            {"x-oss-security-token": security_token},
            {"x-oss-date": dt_obj_1},  
        ]
    }
    print(dt_obj_1)
    policy_str = json.dumps(policy).strip()

    # 步骤2：构造待签名字符串（StringToSign）
    stringToSign = base64.b64encode(policy_str.encode()).decode()

    # 步骤3：计算SigningKey
    dateKey = hmacsha256(("aliyun_v4" + temp_access_key_secret).encode(), dt_obj_2)
    dateRegionKey = hmacsha256(dateKey, "cn-hangzhou")
    dateRegionServiceKey = hmacsha256(dateRegionKey, "oss")
    signingKey = hmacsha256(dateRegionServiceKey, "aliyun_v4_request")

    # 步骤4：计算Signature
    result = hmacsha256(signingKey, stringToSign)
    signature = result.hex()

    # 组织返回数据
    response_data = {
        'policy': stringToSign,  #表单域
        'x_oss_signature_version': "OSS4-HMAC-SHA256",  #指定签名的版本和算法，固定值为OSS4-HMAC-SHA256
        'x_oss_credential': f"{temp_access_key_id}/{dt_obj_2}/cn-hangzhou/oss/aliyun_v4_request",  #指明派生密钥的参数集
        'x_oss_date': dt_obj_1,  #请求的时间
        'signature': signature,  #签名认证描述信息
        'host': host,
        'dir': upload_dir,
        'security_token': security_token  #安全令牌

    }
    return jsonify(response_data)

if __name__ == "__main__":
    app.run(host="0.0.0.0", port=8000)