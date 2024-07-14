import Header from "../Applications/Header";
import { Link } from "react-router-dom";

const Admin = () => {
    return (
        <>
        <Header/>
        <div className="admin-container">
            <h2>管理画面</h2>
            <hr></hr>
            <div className="admin-dashboard-links">
                <Link className="" to={"/admin/users"}>
                    ユーザー
                </Link>
                <Link className="" to={"/admin/qrmarks"}>
                    QRmark
                </Link>
                <Link className="">
                    学校を追加
                </Link>
                <Link className="">
                    ユーザーを追加
                </Link>
                <Link className="">
                    その他
                </Link>
            </div>
        </div>
        </>
    );
}

export default Admin;