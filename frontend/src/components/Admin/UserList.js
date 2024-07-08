import { useState } from "react";
import Header from "../Applications/Header";
import UseFetch from "../Applications/UseFetch";
import moment from "moment/moment";
import { TableContainer, Table, Thead, Tr, Th, Td, Tbody, Icon } from '@chakra-ui/react'
import { Button } from "@chakra-ui/react";
import {ArrowBackIcon, ArrowForwardIcon} from "@chakra-ui/icons"
import { Link } from "react-router-dom";

const UserList = () => {
    const [pageNumber, setPageNumber] = useState(1);
    const {error, isPending, data} = UseFetch(`/user/list?page=${pageNumber}`, [pageNumber]);

    if (error) {
        return (<p>{error.message}</p>);
    }

    const incrementPageNumber = () => {
        setPageNumber(data.page+1);
    }

    const decrementPageNumber = () => {
        setPageNumber(data.page-1);
    }

    moment.locale('ja');

    return (
        <section>
            <Header/>
            <div className="admin-panel-container">
                <div className="admin-panel-middle">
                    <h2>ユーザー</h2>
                    <hr></hr>
                    {data && data.users.length > 1 && 
                    <TableContainer>
                        <Table size={['sm']}>
                            <Thead>
                              <Tr>
                                <Th scope="col">ユーザーID</Th>
                                <Th scope="col">名前</Th>
                                <Th scope="col">メールアドレス</Th>
                                <Th scope="col">権限</Th>
                                <Th scope="col">認証済</Th>
                                <Th scope="col">登録学校</Th>
                                <Th scope="col">時間</Th>
                              </Tr>
                            </Thead>
                            <Tbody>
                                {data.users.map((user, index) => (
                                <Tr key={index}>
                                    <Td>{user.user_id}</Td>
                                    <Td>{user.name}</Td>
                                    <Td>{user.email}</Td>
                                    <Td>{user.role}</Td>
                                    <Td>{user.verified ? "はい" : "いいえ"}</Td>
                                    <Td><Link to={`/school/${user.school.school_id}`}>{user.school.name}</Link></Td>
                                    <Td>{moment(user.created_at).format('YYYY/M/D H:m:s')}</Td>
                                </Tr>
                                ))}
                            </Tbody>
                        </Table>
                    </TableContainer>
                    }
                    {data && (data.page > 1 || data.has_next) && 
                        <div className="pagination">
                          <Button size='sm' isDisabled={isPending || data.page <= 1} onClick={decrementPageNumber}><Icon as={ArrowBackIcon}/></Button>
                          <Button size='sm' isDisabled={isPending || !data.has_next} onClick={incrementPageNumber}><Icon as={ArrowForwardIcon}/></Button>
                        </div>
                    }
                </div>
            </div>
        </section>
    );
}
 
export default UserList;