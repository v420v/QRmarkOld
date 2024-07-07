import Header from "./Applications/Header";
import NotFoundImage from "./NotFound.png"

const NotFound = () => {
    return (
        <>
        <Header/>
        <div className="not-found-page">
            <img src={NotFoundImage}></img>
        </div>
        </>
    );
}
 
export default NotFound;