import { Link, useParams } from "react-router-dom";
import UseFetch from "../Applications/UseFetch";
import moment from "moment/moment";
import Header from "../Applications/Header";
import { TableContainer, Table, Thead, Tr, Th, Td, Tbody, Icon, Button } from '@chakra-ui/react'

const UserDetail = () => {
    moment.locale('ja');
    const { id } = useParams();
    const {error: userDetailError, data: userDetail, } = UseFetch(`/user/${id}`)

    return (
        <>
            <Header/>
            <div className="admin-container">
                <h2>基本情報</h2>
                <hr></hr>
                {
                userDetail && <>
                    <div className="admin-user-detail-name">{userDetail.name}</div>
                    <div className="admin-user-detail-email">{userDetail.email}</div>
                    <div className="admin-user-detail-role">{userDetail.role}</div>
                    <div className="admin-user-detail-verified">{userDetail.verified ? "認証済" : "非認証"}</div>
                    <div className="admin-user-detail-created_at">{moment(userDetail.created_at).format('YYYY/M/D H:m:s')} に作成</div>
                    <div className="admin-user-detail-school"><Link to={`/schools/${userDetail.school.school_id}`}>{userDetail.school.name}</Link> を支援中</div>
                </>
                }
            </div>
        </>
    );
}

export default UserDetail;