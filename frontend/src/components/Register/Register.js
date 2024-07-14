import { useState } from "react";
import axios from '../../api/axios';
import { Link } from "react-router-dom";
import { useDebounce } from 'react-use';
import { FontAwesomeIcon } from "@fortawesome/react-fontawesome";
import { faChevronLeft } from "@fortawesome/free-solid-svg-icons";
import { faChevronRight } from "@fortawesome/free-solid-svg-icons";
import { FormControl, FormLabel, Input, FormHelperText, Button, useDisclosure, Modal, ModalOverlay, ModalContent, ModalBody,
} from "@chakra-ui/react"
import PasswordInput from "../Applications/Password"
import Header from "../Applications/Header";

const SchoolList = ({setSchoolName, onClose, schoolID, setSchoolID}) => {
    const [pageNumber, setPageNumber] = useState(1);
    const [schools, setSchools] = useState(null);
    const [isPending, setIsPending] = useState(false);
    const [error, setError] = useState(null);
    const [schoolQuery, setSchoolQuery] = useState('');

    useDebounce(() => {
        const handleFunc = async () => {
            setIsPending(true);
            if (schoolQuery === '') {
                setSchoolQuery([]);
                setIsPending(false);
                return
            }
            let ignore = false;
            try {
                const response = await axios.get(`/school/search?q=${schoolQuery}&page=${pageNumber}`, {
                    headers: { 'Content-Type': 'text/plain; charset=utf-8' }
                });
                if (!ignore) {
                    setSchools(response.data);
                }
            } catch (err) {
              setError(err);
            } finally {
                setIsPending(false);
                return () => {
                    ignore = true;
                };
            }
        }
        handleFunc();
    }, 500, [schoolQuery, pageNumber]);

    if (error) {
        return <p>{error.message}</p>
    }

    const incrementPageNumber = () => {
        setPageNumber(schools.page+1)
    }

    const decrementPageNumber = () => {
        setPageNumber(schools.page-1)
    }

    return (
        <section>
            <Input type="text" placeholder="支援したい学校名を検索" onChange={(e) => {setSchoolQuery(e.target.value)}} value={schoolQuery}/>
            {(schools && schools.schools.length > 0) &&
                <>
                <div className="select-school-list">
                    {schools.schools.map((school) => (
                    school.school_id !== schoolID ? 
                        <div key={school.school_id} className="select-school-list-school" value={school.name} onClick={() => {setSchoolID(school.school_id); setSchoolName(school.name); onClose();}}>
                            {school.name}
                        </div>
                    :
                        <div key={school.school_id} className="select-school-list-school select-school-list-school-selected" value={school.name} onClick={() => {setSchoolID(school.school_id); setSchoolName(school.name); onClose();}}>
                            {school.name}
                        </div>
                    ))}
                </div>
                <div className="pagination">
                    <Button isDisabled={isPending || schools.page <= 1} onClick={decrementPageNumber}><FontAwesomeIcon icon={faChevronLeft}></FontAwesomeIcon></Button>
                    <Button isDisabled={isPending || !schools.has_next} onClick={incrementPageNumber}><FontAwesomeIcon icon={faChevronRight}></FontAwesomeIcon></Button>
                </div>
                </>
            }
        </section>
    );
}

const Register = () => {
    const [name, setName] = useState('');
    const [email, setEmail] = useState('');
    const [pwd, setPwd] = useState('');
    const [schoolID, setSchoolID] = useState(0);
    const [schoolName, setSchoolName] = useState(0);
    const [registerIsPending, setRegisterIsPending] = useState(false);
    const {isOpen, onOpen, onClose } = useDisclosure();

    const [registerSuccess, setRegisterSuccess] = useState(false);

    const [errMsg, setErrMsg] = useState('');

    const [verifyingEmail, setVerifyingEmail] = useState("");

    const handleSubmit = async (e) => {
        e.preventDefault();
        setRegisterIsPending(true);

        try {
            await axios.post('/user',
                JSON.stringify({
                    "name": name,
                    "email": email,
                    "password": pwd,
                    "school_id": schoolID
                }),
                {
                    headers: { 'Content-Type': 'text/plain; charset=utf-8' }
                     
                }
            );
            setVerifyingEmail(email);
            setEmail('');
            setPwd('');
            setSchoolID(0);
            setName('');
            setRegisterSuccess(true);
            setRegisterIsPending(false);
        } catch (err) {
            setRegisterIsPending(false);
            if (!err?.response) {
                setErrMsg('サーバーの応答がありません');
            } else {
                setErrMsg('アカウント作成に失敗')
            }
        }
    }

    if (registerSuccess) {
        return (
            <>
            <div className="register-container">
                <div className="register-success-message">
                    <h2>メールを検証する</h2>
                    <br></br>
                    <p>{verifyingEmail}宛に認証リンクを送りました。</p>
                </div>
            </div>
            </>
        )
    }

    return (
        <>
        <div className="register-container">
            <h2>新規登録</h2>

            <Modal isOpen={isOpen} onClose={onClose}>
              <ModalOverlay />
              <ModalContent>
                <ModalBody>
                    <SchoolList setSchoolName={setSchoolName} onClose={onClose} schoolID={schoolID} setSchoolID={setSchoolID}/>
                </ModalBody>
              </ModalContent>
            </Modal>

            {errMsg && errMsg}
            <form onSubmit={handleSubmit}>
                <FormControl>
                    <FormLabel mt="15px">名前</FormLabel>
                    <Input id="name" placeholder="名前" type='name' onChange={(e) => setName(e.target.value)} value={name} />

                    <FormLabel mt="15px">メールアドレス</FormLabel>
                    <Input id="email" placeholder="メールアドレス" type='email' onChange={(e) => setEmail(e.target.value)} value={email} />

                    <FormLabel mt="15px">パスワード</FormLabel>
                    <PasswordInput pwd={pwd} setPwd={setPwd}/>

                    <FormLabel mt="15px">学校を選択</FormLabel>
                    <Button onClick={onOpen} width="100%" size="sm">{schoolName ? `${schoolName}` : "学校を選択する"}</Button>
                    <FormHelperText>支援する学校は後で変更可能 (WIP)</FormHelperText>

                    <Button isLoading={registerIsPending} type={'submit'}  mt="15px" width="100%" color="#fff" bg="#e5ad15" _hover={{ bg: '#e5ae15d0'}}>登録</Button>
                </FormControl>
            </form>
            <p className="register-form-bottom-text">
                登録済の方は<Link to="/login" replace>こちら</Link>からログイン
            </p>
        </div>
        </>
    );
}

export default Register