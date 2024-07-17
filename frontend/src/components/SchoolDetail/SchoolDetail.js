import { useParams } from "react-router-dom";
import Header from "../Applications/Header";
import UseFetch from "../Applications/UseFetch";
import * as XLSX from 'xlsx';
import { useState } from "react";
import { TableContainer, Table, Thead, Tr, Th, Td, Tbody, Button } from '@chakra-ui/react'

const SchoolDetail = () => {
    const { id } = useParams();

    const {error: schoolDetailError, data: schoolDetail, } = UseFetch(`/schools/${id}`)
    const {error: schoolPointsError, isPending: schoolPointsIsPending, data: schoolPoints} = UseFetch(`/schools/${id}/points`);

    const [downloadIsPending, setDownloadIsPending] = useState(false);

    const downloadExcel = () => {
        setDownloadIsPending(true);

        const table = document.getElementById('table');

        const ws = XLSX.utils.table_to_sheet(table);
    
        const wb = XLSX.utils.book_new();
        XLSX.utils.book_append_sheet(wb, ws, 'Sheet1');
    
        XLSX.writeFile(wb, `${schoolDetail.name}.xlsx`);

        setDownloadIsPending(false);
    };

    if (schoolPointsError) {
        return <>{schoolPointsError.message}</>
    }

    if (schoolDetailError) {
        return <>{schoolDetailError.message}</>
    }

    return (
        <section>
            <Header/>
            <div className="school-detail-container">
                <div className="school-detail-middle">
                    <h2>{schoolDetail && schoolDetail.name}</h2>
                    <hr></hr>
                    {schoolPoints && schoolPoints.length > 1 && 
                    <TableContainer>
                        <Table id="table" size={['sm']}>
                            <Thead>
                              <Tr>
                                <Th id="th-company">企業名</Th>
                                <Th isNumeric id="th-point">ポイント</Th>
                              </Tr>
                            </Thead>
                            <Tbody>
                                {schoolPoints.map((ssp, index) => (
                                    <Tr key={index}>
                                        <Td className="td-company-name">{ssp.company.name}</Td>
                                        <Td className="td-point" isNumeric>{ssp.points}</Td>
                                    </Tr>
                                ))}
                            </Tbody>
                        </Table>
                    </TableContainer>
                    }
                    <div className="school-detail-table-top">
                        <Button size="sm" isDisabled={schoolPointsIsPending || downloadIsPending} onClick={downloadExcel} >ダウンロード</Button>
                    </div>
                </div>
            </div>
        </section>
    );
}
 
export default SchoolDetail;

