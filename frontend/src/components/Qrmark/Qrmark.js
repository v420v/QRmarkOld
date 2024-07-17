import { QrReader } from '@cmdnio/react-qr-reader';

import { useNavigate, useLocation } from 'react-router-dom';
import React, { useState } from 'react';
import axios from '../../api/axios';
import Header from '../Applications/Header';

const Qrmark = () => {
    const [isPending, setIsPending] = useState(false);
    const [error, setError] = useState(null);
    const [cameraActive, setCameraActive] = useState(true);
    const navigate = useNavigate();
    let location = useLocation();
    const from = location.state?.form?.pathname || "/";

    const handleRetry = () => {
        setError(null);
        setCameraActive(true);
    };

    return (
        <section>
            <Header/>
            <div className="qrmark-reader-container">
                <div className="qrmark-reader">
                    {cameraActive ? (
                        <QrReader
                            constraints={{
                                facingMode: 'environment'
                            }}
                            onResult={(result) => {
                                if ((!!result) && !isPending && !error) {
                                    setIsPending(true);

                                    const handleSubmit = async () => {
                                        setIsPending(true);
                                        if (isPending) {
                                            return
                                        }
                                        try {
                                            await axios.post(`/qrmarks`,
                                                JSON.stringify({
                                                    "jwt": result.text
                                                }),
                                                {
                                                    headers: { 'Content-Type': 'text/plain; charset=utf-8' }
                                                     
                                                }
                                            );

                                            setIsPending(false);
                                            navigate(from, { replace: true });
                                        } catch (err) {
                                            if (err.response && err.response.status === 401) {
                                                navigate("/login", { replace: true });
                                            } else {
                                                setError("エラーが発生しました。もう一度お試しください。");
                                                setCameraActive(false);
                                            }
                                            setIsPending(false);
                                        }
                                    }

                                    handleSubmit();
                                }
                            }}
                        />
                    ) : <div>カメラ停止中</div>}
                </div>
                {error && (
                    <div className="qrmark-reader-error-message">
                        <p>{error}</p>
                        <a onClick={handleRetry}>読み取りをし直す</a>
                    </div>
                )}
            </div>
        </section>
    );
};
 
export default Qrmark;
