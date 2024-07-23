/**
 * 登录页面
 *
 * @author jixiangup
 * @since 1.0.0
 */
import {Container} from "@/components/sharedstyles";
import {Button, FormControl, FormLabel, Input} from "@mui/joy";

export default function LoginPage() {
  return (
    <Container>
      <form onSubmit={async (event) => {event.defaultPrevented}}>
        <FormControl required>
          <FormLabel>用户名</FormLabel>
          <Input name='username' />
        </FormControl>
        <Button>登录</Button>
      </form>
    </Container>
  );
}