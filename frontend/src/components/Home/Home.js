import { useContext } from "react";
import { useState } from "react";
import Moment from "moment";
import Header from "../Applications/Header";
import { Link } from "react-router-dom";
import UseFetch from "../Applications/UseFetch";
import CurrentUserContext from "../../context/AuthProvider";
import { FontAwesomeIcon } from '@fortawesome/react-fontawesome';
import moment from "moment/moment";
import { faChevronLeft } from "@fortawesome/free-solid-svg-icons";
import { faChevronRight } from "@fortawesome/free-solid-svg-icons";
import { Button, TableContainer, Table, Thead, Tr, Th, Td, Tbody, Stack, Skeleton } from '@chakra-ui/react'
import { useNavigate } from 'react-router-dom';

const TotalPoints = () => {
  const [CurrentUser] = useContext(CurrentUserContext);
  const {error: totalPointsError, totalPointsIsPending, data: totalPointsData} = UseFetch(`/user/${CurrentUser.user_id}/total_points`)
  if (totalPointsError) {
    return <>読み込みに失敗</>;
  }
  if (totalPointsIsPending) {
    return <>読み込み中</>;
  }

  if (totalPointsData) {
    return <>{totalPointsData.points}</>;
  }
}

const Home = () => {
  const [pageNumber, setPageNumber] = useState(1);
  const [CurrentUser, ] = useContext(CurrentUserContext);
  const {error: historyError, isPending: historyIsPending, data: historyData} = UseFetch(`/user/${CurrentUser.user_id}/qrmark/list?page=${pageNumber}`, [pageNumber])
  const navigate = useNavigate();
  Moment.locale('ja');

  if (historyError) {
    return <>{historyError.message}</>
  }

  const incrementPageNumber = () => {
    setPageNumber(historyData.page+1)
  }

  const decrementPageNumber = () => {
    setPageNumber(historyData.page-1)
  }

  return (
    <>
    <Header/>
    <div className="home-container">
      <div className="home-left">
        <Link className="">
          よくある質問
        </Link>
        <Link className="">
          利用規約
        </Link>
        <Link className="">
          お問い合わせ
        </Link>
      </div>

      <div className="home-middle">
        <h2>ホーム</h2>
        <hr></hr>
        <div className="">寄付額 <span className="home-total-points"><TotalPoints/></span> 円</div>
        <Button width="100%" color="#fff" bg="#e5ad15" _hover={{ bg: '#e5ae15d0'}} onClick={() => {navigate("/qrmark");}}>QRmarkを使用</Button>
        {(historyData && historyData.qrmarks.length > 0) &&
          <div className="home-qrmark-history">
            <TableContainer>
              <Table size={['sm', 'md']}>
                <Thead>
                  <Tr>
                    <Th>日付</Th>
                    <Th>企業名</Th>
                    <Th isNumeric>ポイント</Th>
                  </Tr>
                </Thead>
                <Tbody>
                  {historyData.qrmarks.map((history, index) => (
                  <Tr key={index}>
                    <Td>{moment(history.created_at).format('M/D')}</Td>
                    <Td>{history.company_name}</Td>
                    <Td isNumeric color="#55b958">+ {history.points}円</Td>
                  </Tr>
                  ))}
                </Tbody>
              </Table>
            </TableContainer>
          </div>
        }
        {historyData && (historyData.page > 1 || historyData.has_next) && 
          <div className="pagination">
            <Button isDisabled={historyData.page <= 1 || historyIsPending} onClick={decrementPageNumber}><FontAwesomeIcon icon={faChevronLeft}></FontAwesomeIcon></Button>
            <Button isDisabled={!historyData.has_next || historyIsPending} onClick={incrementPageNumber}><FontAwesomeIcon icon={faChevronRight}></FontAwesomeIcon></Button>
          </div>
        }
      </div>
    </div>
    </>
  );
};

export default Home;