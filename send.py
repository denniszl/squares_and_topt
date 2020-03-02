import requests
import base64
from base64 import b64encode
import pyotp
import hmac
import hashlib
import time
import struct
from requests.auth import HTTPBasicAuth
from requests_toolbelt.utils import dump

def get_hotp_token(intervals_no):
    encodedKey = base64.b32encode('*')
    decodedKey = base64.b32decode(encodedKey)
    msg = struct.pack(">Q", intervals_no)
    # msg = str(intervals_no)
    h = hmac.new(decodedKey, msg, hashlib.sha512).digest()
    # return h
    o = ord(h[-1]) & 15
    h = (struct.unpack(">L", h[o:o+4])[0] & 0x7fffffff) % 10000000000
    strH = str(h)
    while len(strH) < 10:
        print 'appending'
        strH = "0" + strH
    return strH

def get_totp_token():
    return get_hotp_token(int(time.time())//30)


def hotp(key, counter, digits=10, digest='sha512'):
    key = base64.b32decode(key.upper() + '=' * ((8 - len(key)) % 8))
    counter = struct.pack('>Q', counter)
    mac = hmac.new(key, counter, hashlib.sha512).digest()
    offset = ord(mac[-1]) & 0x0f
    binary = struct.unpack('>L', mac[offset:offset+4])[0] & 0x7fffffff
    return str(binary)[-digits:].rjust(digits, '0')


def totp(key, time_step=30, digits=10, digest='sha512'):
    return hotp(key, int(time.time() / time_step), digits, digest)

url = 'https://api.challenge.*.com/challenges/003'
dataToPost = {
  "github_url": "*",
  "contact_email": "*"
}

h = {
    'Content-Type': 'application/json'
}

password = get_totp_token()
print password
# print 'basic auth:', HTTPBasicAuth('denniszl14@gmail.com', password)

# print topt.now()
resp = requests.post(url, json=dataToPost, headers=h, auth=('*', password))
# data = dump.dump_all(resp)
# print(data.decode('utf-8'))
if resp.status_code == 200:
    print 'done'
    exit(0)
else:
    print resp.status_code
    print resp.json()
    # import ipdb
    # ipdb.set_trace()
# import ipdb; ipdb.set_trace()