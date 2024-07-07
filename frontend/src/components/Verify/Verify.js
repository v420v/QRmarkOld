import { useParams } from "react-router-dom";
import UseFetch from "../Applications/UseFetch";
import Header from "../Applications/Header";
import { Spinner } from '@chakra-ui/react'

const Verify = () => {
    const { token } = useParams();

    const {error, isPending, } = UseFetch(`/verify/${token}`);

    return (
        <>
            <Header/>
            <div className="verify-container">
                {error ?
                    <h2>認証に失敗しました。</h2>
                    : isPending ?
                    <><Spinner size={['md', 'lg']} /></>
                    :
                    <div className="verify-success-message">
                        <h2>認証完了</h2>
                    </div>
                }
            </div>
        </>
    );
}
 
export default Verify;