import { Link } from 'react-router-dom';
import { useContext } from "react";
import CurrentUserContext from "../../context/AuthProvider";
import { Menu, MenuButton, IconButton, MenuList, MenuItem, Button } from '@chakra-ui/react';
import { useNavigate } from 'react-router-dom';

import {HamburgerIcon} from "@chakra-ui/icons"

const Header = () => {
    const [CurrentUser, setCurrentUser] = useContext(CurrentUserContext);

    const navigate = useNavigate();

    const handleSubmit = async () => {
      setCurrentUser(null);
    }

    return (
        <header>
            {CurrentUser ?
            <Menu>
              <MenuButton
                as={IconButton}
                aria-label='Options'
                icon={<HamburgerIcon />}
                variant='outline'
              />
              <MenuList>
                <Link to={`/`}>
                    <MenuItem>
                      ホーム
                    </MenuItem>
                </Link>
                <Link to={`/school/${CurrentUser.school_id}`}>
                    <MenuItem>
                      支援してる学校
                    </MenuItem>
                </Link>
                <Link to="/" replace={true} onClick={handleSubmit}>
                    <MenuItem>
                      ログアウト
                    </MenuItem>
                </Link>
                {CurrentUser.role === "admin" &&
                <Link to="/admin">
                    <MenuItem>
                        管理者
                    </MenuItem>
                </Link>}
              </MenuList>
            </Menu>
            :
            <Button color="#fff" bg="#e5ad15" _hover={{ bg: '#e5ae15d0'}} onClick={() => {navigate("/login");}}>ログイン</Button>
            }
        </header>
    );
}

export default Header;