import { useState } from "react";
import Header from "../Applications/Header";
import UseFetch from "../Applications/UseFetch";
import moment from "moment/moment";
import { TableContainer, Table, Thead, Tr, Th, Td, Tbody, Icon } from '@chakra-ui/react'
import { Button } from "@chakra-ui/react";
import {ArrowBackIcon, ArrowForwardIcon} from "@chakra-ui/icons"

const Admin = () => {
    const [pageNumber, setPageNumber] = useState(1);
    const {error, isPending, data} = UseFetch(`/qrmark/list?page=${pageNumber}`, [pageNumber]);

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
                    <h2>履歴</h2>
                    <hr></hr>
                    {data && data.qrmarks.length > 1 && 
                    <TableContainer>
                        <Table size={['sm', 'md']}>
                            <Thead>
                              <Tr>
                                <Th>QRMARK_ID</Th>
                                <Th scope="col">ユーザーID</Th>
                                <Th scope="col">学校</Th>
                                <Th scope="col">企業</Th>
                                <Th scope="col">ポイント</Th>
                                <Th scope="col">時間</Th>
                              </Tr>
                            </Thead>
                            <Tbody>
                                {data.qrmarks.map((qrmark, index) => (
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
                    {data && data.page > 1 && 
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
 
export default Admin;