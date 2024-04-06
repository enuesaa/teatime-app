import { Header } from '@/components/common/Header'
import { ListTeas } from '@/components/top/ListTeas'
import { Container } from '@radix-ui/themes'
import { Link } from 'react-router-dom'

export default function Page() {
  return (
    <>
      <Header />
      <Container size='4' m='3' pt='3'>
        <Link to='/teapods/links/teas/links'>links</Link>
      </Container>
    </>
  )
}
