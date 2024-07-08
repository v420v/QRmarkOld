import Header from "../Applications/Header";
import UseFetch from "../Applications/UseFetch";
import moment from "moment/moment";
import { TableContainer, Table, Thead, Tr, Th, Td, Tbody, Stack, Skeleton } from '@chakra-ui/react'
import { Card, CardHeader, CardBody, CardFooter, Heading, Text } from "@chakra-ui/react";
import { Link } from "react-router-dom";

const Admin = () => {
    const {error: QRmarkHistoryError, isPending: QRmarkHistoryIsPending, data: QRmarkHistoryData} = UseFetch(`/qrmark/list?page=1`);
    const {error: UserListError, isPending: UserListIsPending, data: UserListData} = UseFetch(`/user/list?page=1`);

    return (
        <>
        <Header/>
        <div className="admin-dashboard-container">
            <div className="admin-dashboard-top">
                <h2>管理画面</h2>
                <hr></hr>
            </div>
            <div className="admin-dashboard">
                <div className="card admin-dashboard-left">
                    <Card variant={['outline']}>
                        <CardHeader>
                            <Heading size='md'>ユーザー</Heading>
                        </CardHeader>
                        <CardBody>
                        {UserListError ? UserListError.message :
                            UserListData && <>
                            <TableContainer>
                                <Table size={['sm']}>
                                    <Thead>
                                      <Tr>
                                        <Th scope="col">名前</Th>
                                        <Th scope="col">時間</Th>
                                      </Tr>
                                    </Thead>
                                    <Tbody>
                                        {UserListData.users.slice(0, 5).map((user, index) => (
                                        <Tr key={index}>
                                            <Td>{user.name}</Td>
                                            <Td>{moment(user.created_at).format('YYYY/M/D')}</Td>
                                        </Tr>
                                        ))}
                                    </Tbody>
                                </Table>
                            </TableContainer>
                            </>}
                        </CardBody>
                        <CardFooter>
                            <Link to="/admin/users">すべて見る</Link>
                        </CardFooter>
                    </Card>
                </div>
                <div className="card admin-dashboard-middle">
                    <Card className="card" variant={['outline']}>
                        <CardHeader>
                            <Heading size='md'>QRmark</Heading>
                        </CardHeader>
                        <CardBody>
                            {QRmarkHistoryError ? QRmarkHistoryError.message :
                            QRmarkHistoryData && <>
                            <TableContainer>
                                <Table size={['sm']}>
                                    <Thead>
                                      <Tr>
                                        <Th scope="col">ユーザーID</Th>
                                        <Th scope="col">ポイント</Th>
                                        <Th scope="col">時間</Th>
                                      </Tr>
                                    </Thead>
                                    <Tbody>
                                        {QRmarkHistoryData.qrmarks.slice(0, 5).map((qrmark, index) => (
                                        <Tr key={index}>
                                            <Td>{qrmark.user_id}</Td>
                                            <Td>{qrmark.points}</Td>
                                            <Td>{moment(qrmark.created_at).format('YYYY/M/D H:m:s')}</Td>
                                        </Tr>
                                        ))}
                                    </Tbody>
                                </Table>
                            </TableContainer>
                            </>}
                        </CardBody>
                        <CardFooter>
                            <Link to="/admin/qrmarks">すべて見る</Link>
                        </CardFooter>
                    </Card>
                </div>
            </div>
        </div>
        </>
    );
}
 
export default Admin;