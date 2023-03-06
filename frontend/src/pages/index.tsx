import { Header } from '@/components/common/Header'
import { Main } from '@/components/common/Main'
import { LinkCard } from '@/components/common/LinkCard'

export default function TopPage() {
  return (
    <>
      <Header />
      <Main>
        <LinkCard href='/lounge' text='lounge' />
      </Main>
    </>
  )
}
