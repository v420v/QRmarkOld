import { Link, useParams } from "react-router-dom";
import UseFetch from "../Applications/UseFetch";
import moment from "moment/moment";
import Header from "../Applications/Header";
import { useState } from "react";
import { TableContainer, Table, Thead, Tr, Th, Td, Tbody, Icon, Button } from '@chakra-ui/react'
import {ArrowBackIcon, ArrowForwardIcon} from "@chakra-ui/icons"
import { MdAdminPanelSettings } from "react-icons/md";
import { MdVerified } from "react-icons/md";

const UserDetail = () => {
    moment.locale('ja');
    const { id } = useParams();
    const {error: userDetailError, data: userDetail, } = UseFetch(`/user/${id}`)

    const [pageNumber, setPageNumber] = useState(1);
    const {error: historyError, isPending: historyIsPending, data: historyData} = UseFetch(`/qrmark/list?page=${pageNumber}&user=${id}`, [pageNumber]);

    if (userDetailError) {
        return <>{userDetailError.message}</>
    }

    if (historyError) {
        return <>{historyError.message}</>
    }

    const incrementPageNumber = () => {
        setPageNumber(historyData.page+1);
    }

    const decrementPageNumber = () => {
        setPageNumber(historyData.page-1);
    }

    return (
        <>
            <Header/>
            <div className="admin-container">
                <h2>基本情報</h2>
                <hr></hr>
                {
                userDetail && <>
                    <div className="admin-user-detail-name" style={{display: "flex", alignItems: "center", gap: "5px"}}>
                        {userDetail.name}
                        {userDetail.role === "admin" && <MdAdminPanelSettings color="#e5ad15" fill="#e5ad15" />}
                        {userDetail.verified && <MdVerified color="#1DA1F2" fill="#1DA1F2" />}
                    </div>
                    <div className="admin-user-detail-email">{userDetail.email}</div>
                    <div className="admin-user-detail-created_at">{moment(userDetail.created_at).format('YYYY/M/D H:m:s')} に作成</div>
                    <div className="admin-user-detail-school"><Link to={`/school/${userDetail.school.school_id}`}>{userDetail.school.name}</Link></div>
                </>
                }
                <br></br>
                <h2>QRmark情報</h2>
                <hr></hr>
                {historyData && historyData.qrmarks.length > 0 && 
                <TableContainer>
                    <Table size={['sm']}>
                        <Thead>
                            <Tr>
                                <Th scope="col">ID</Th>
                                <Th scope="col">ユーザーID</Th>
                                <Th scope="col">学校名</Th>
                                <Th scope="col">企業</Th>
                                <Th scope="col">ポイント</Th>
                                <Th scope="col">時間</Th>
                            </Tr>
                        </Thead>
                        <Tbody>
                            {historyData.qrmarks.map((qrmark, index) => (
                            <Tr key={index}>
                                <Td>{qrmark.qrmark_id}</Td>
                                <Td>{qrmark.user_id}</Td>
                                <Td>{qrmark.school_name}</Td>
                                <Td>{qrmark.company_name}</Td>
                                <Td>{qrmark.points}</Td>
                                <Td>{moment(qrmark.created_at).format('YYYY/M/D H:m:s')}</Td>
                            </Tr>
                            ))}
                        </Tbody>
                    </Table>
                </TableContainer>
                }
                {historyData && (historyData.page > 1 || historyData.has_next) && 
                    <div className="pagination">
                      <Button size='sm' isDisabled={historyIsPending || historyData.page <= 1} onClick={decrementPageNumber}><Icon as={ArrowBackIcon}/></Button>
                      <Button size='sm' isDisabled={historyIsPending || !historyData.has_next} onClick={incrementPageNumber}><Icon as={ArrowForwardIcon}/></Button>
                    </div>
                }
                </div>
        </>
    );
}

export default UserDetail;