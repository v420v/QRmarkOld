import { Link } from 'react-router-dom';
import { useContext } from "react";
import CurrentUserContext from "../../context/AuthProvider";
import { Menu, MenuButton, IconButton, MenuList, MenuItem, Button } from '@chakra-ui/react';
import { useNavigate } from 'react-router-dom';
import {HamburgerIcon} from "@chakra-ui/icons"
import { useState } from 'react';
import axios from "../../api/axios";

const Header = () => {
  const [CurrentUser, setCurrentUser] = useContext(CurrentUserContext);

  const navigate = useNavigate();

  const [logoutIsPending, setLogoutIsPending] = useState(false);

  const handleLogout = async (e) => {
    setLogoutIsPending(true);
    e.preventDefault();

    try {
      await axios.delete("/logout",
        JSON.stringify({}),
        {
          headers: { 'Content-Type': 'text/plain; charset=utf-8' },
          withCredentials: true,
        }
      );
      setLogoutIsPending(false);
      setCurrentUser(null);
    } catch (err) {
      setLogoutIsPending(false);
    }
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
                <Button _hover={{ textDecoration: "none" }} color={"#464956"} width={"100%"} isDisabled={logoutIsPending} onClick={handleLogout} variant='link'>
                  <MenuItem>
                    ログアウト
                  </MenuItem>
                </Button>
                {CurrentUser.role === "admin" &&
                <Link to="/admin">
                    <MenuItem>
                        管理者
                    </MenuItem>
                </Link>}
              </MenuList>
            </Menu>
            :
              <>
                <Button color="#fff" bg="#e5ad15" _hover={{ bg: '#e5ae15d0'}} onClick={() => {navigate("/login");}}>ログイン</Button>
              </>
            }
        </header>
    );
}

export default Header;