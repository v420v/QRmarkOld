import { useState } from 'react';
import { useNavigate, useLocation, Link } from 'react-router-dom';
import axios from '../../api/axios';
import PasswordInput from "../Applications/Password"
import {FormControl, FormLabel, Input, Button} from "@chakra-ui/react"
import Header from '../Applications/Header';

const Login = () => {
    const navigate = useNavigate();
    const location = useLocation();
    const from = location.state?.form?.pathname || "/";

    const [email, setEmail] = useState('');
    const [pwd, setPwd] = useState('');
    const [errMsg, setErrMsg] = useState('');

    const [loginIsPending, setLoginIsPending] = useState(false);

    const handleSubmit = async (e) => {
        setLoginIsPending(true);
        e.preventDefault();

        try {
            await axios.post("/login",
                JSON.stringify({
                    "email": email,
                    "password": pwd
                }),
                {
                    headers: { 'Content-Type': 'text/plain; charset=utf-8' },
                    withCredentials: true,
                }
            );
            setEmail('');
            setPwd('');

            setLoginIsPending(false);
            navigate(from, {replace: true});
        } catch (err) {
            setLoginIsPending(false);
            if (!err?.response) {
                setErrMsg('サーバーからの応答なし');
            } else {
                setErrMsg('ログインに失敗');
            }
        }
    }

    return (
        <>
        <Header/>
        <div className="register-container">
            <h2>ログイン</h2>
            {errMsg && errMsg}
            <form onSubmit={handleSubmit}>
                <FormControl>
                    <FormLabel mt="15px">メールアドレス</FormLabel>
                    <Input id="name" placeholder="メールアドレス" type='email' onChange={(e) => setEmail(e.target.value)} value={email} />

                    <FormLabel mt="15px">パスワード</FormLabel>
                    <PasswordInput pwd={pwd} setPwd={setPwd}/>

                    <Button isLoading={loginIsPending} isDisabled={loginIsPending} type="submit" mt="15px" width="100%" color="#fff" bg="#e5ad15" _hover={{ bg: '#e5ae15d0'}}>登録</Button>
                </FormControl>
            </form>
            <p className="register-form-bottom-text">
            アカウントをお持ちでない方は<Link to="/Register" replace>こちら</Link>
            </p>
        </div>
        </>
    )
}

export default Login