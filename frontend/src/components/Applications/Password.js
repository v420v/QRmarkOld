import { Input, InputGroup, InputRightElement, Button } from "@chakra-ui/react"
import { useState } from "react"

const PasswordInput = ({pwd, setPwd}) => {
    const [show, setShow] = useState(false)
    const handleClick = () => setShow(!show)
  
    return (
      <InputGroup size='md'>
        <Input
          pr='4.5rem'
          type={show ? 'text' : 'password'}
          id="password"
          name="password"
          placeholder='Enter password'
          onChange={(e) => setPwd(e.target.value)} value={pwd}
        />
        <InputRightElement width='4.5rem'>
          <Button h='1.75rem' size='sm' onClick={handleClick}>
            {show ? '表示' : '非表示'}
          </Button>
        </InputRightElement>
      </InputGroup>
    )
}
 
export default PasswordInput;