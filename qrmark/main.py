import jwt
import qrcode
from qrcode.image.styledpil import StyledPilImage
from qrcode.image.styles.moduledrawers.pil import RoundedModuleDrawer, VerticalBarsDrawer, GappedSquareModuleDrawer
from qrcode.image.styles.colormasks import SolidFillColorMask
from jwt import InvalidSignatureError
from cryptography.hazmat.backends import default_backend
from cryptography.hazmat.primitives import serialization
from cryptography.hazmat.primitives.asymmetric import ec
from typing import List, Tuple, Union
import random
import os

def load_ecdsa_secret_key_from_file(file_path: str) -> Union[ec.EllipticCurvePrivateKey, None]:
    try:
        with open(file_path, 'rb') as key_file:
            pem_data = key_file.read()
            private_key = serialization.load_pem_private_key(
                pem_data,
                password=None,
                backend=default_backend()
            )
            if not isinstance(private_key, ec.EllipticCurvePrivateKey):
                raise ValueError("これはEC秘密鍵ではありません")
            return private_key
    except Exception as e:
        print(f"EC秘密鍵の読み込みに失敗しました: {e}")
        return None

qrmark_number = [
    1,2,3,4,5,6,7,8,9,10,11,12,13,14,15,16,17,18,19,20,21,22,23,24,25,26,27,28,29,30
]

def generate_jwt_with_ecdsa_signature(private_key: ec.EllipticCurvePrivateKey, qrmark_id: int, n: int) -> Tuple[str, Union[None, Exception]]:
    payload = {
        "sub":   qrmark_id,
		"num":   n,
		"point": random.randint(1, 3),
    }

    try:
        token = jwt.encode(payload, private_key, algorithm='ES256')
        return token, None
    except InvalidSignatureError as e:
        return '', e

def main():
    if not os.path.isdir('images'):
        os.mkdir("images")

    private_key = load_ecdsa_secret_key_from_file("../keys/ecdsa_p256_private_key.pem")
    if private_key is None:
        return

    id = 301

    for n in qrmark_number:
        if not os.path.isdir(f"images/qrmark-number-{n}"):
            os.mkdir(f"images/qrmark-number-{n}")
        for i in range(30):
            token, err = generate_jwt_with_ecdsa_signature(private_key, id, n)

            if err:
                print(f"JWT生成に失敗しました: {err}")
                return

            qr = qrcode.QRCode(
                error_correction=qrcode.constants.ERROR_CORRECT_H
            )
            print(token)
            qr.add_data(token)
            img = qr.make_image(
                image_factory=StyledPilImage,
                module_drawer=RoundedModuleDrawer(),
                eye_drawer=RoundedModuleDrawer(),
                color_mask=SolidFillColorMask(front_color=(0, 0, 0))
            )
            img.save(f"images/qrmark-number-{n}/qrmark-{n}-{i}-{id}.png")

            id += 1

if __name__=='__main__':
    main()

