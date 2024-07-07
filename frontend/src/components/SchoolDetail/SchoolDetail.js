import { useParams } from "react-router-dom";
import Header from "../Applications/Header";
import UseFetch from "../Applications/UseFetch";
import * as XLSX from 'xlsx';
import { useState } from "react";
import { TableContainer, Table, Thead, Tr, Th, Td, Tbody, Input, Button } from '@chakra-ui/react'

const SchoolDetail = () => {
    const { id } = useParams();

    const getCurrentYearMonth = () => {
        const now = new Date();
        const year = now.getFullYear();
        const month = String(now.getMonth() + 1).padStart(2, '0');
        return `${year}-${month}`;
    };

    const [date, setDate] = useState(getCurrentYearMonth());

    const handleDateChange = (event) => {
        setDate(event.target.value);
    };

    const extractYearAndMonth = (date) => {
        if (date) {
            const [year, month] = date.split('-');
            return { year, month };
        }
        return { year: '', month: '' };
    };

    const { year, month } = extractYearAndMonth(date);

    const {error: schoolDetailError, data: schoolDetail, } = UseFetch(`/school/${id}`)
    const {error: schoolPointsError, isPending: schoolPointsIsPending, data: schoolPoints} = UseFetch(`/school/${id}/points?month=${month}&year=${year}`, [month, year]);

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
                    <Input type="month" value={date} onChange={handleDateChange} />
                    {schoolPoints && schoolPoints.length > 1 && 
                    <TableContainer>
                        <Table size={['sm', 'md', 'lg']}>
                            <Thead>
                              <Tr>
                                <Th>企業名</Th>
                                <Th isNumeric>ポイント</Th>
                              </Tr>
                            </Thead>
                            <Tbody>
                                {schoolPoints.map((ssp, index) => (
                                    <Tr key={index}>
                                        <Td>{ssp.company.name}</Td>
                                        <Td isNumeric>{ssp.points}</Td>
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

