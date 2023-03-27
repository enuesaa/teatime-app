import { PageTitle } from '@/components/common/PageTitle'
import { css, useTheme } from '@emotion/react'
import Link from 'next/link'
import { FaPlus } from 'react-icons/fa'
import { ConfigureFeeds } from '@/components/feed/ConfigureFeeds'

export const Configure = () => {
  const theme = useTheme()

  const styles = {
    main: css({
      margin: '20px',
      padding: '0 10px 10px 10px',
      color: theme.color.main,
    }),
  }

  return (
    <section css={styles.main}>
      <PageTitle title='Feed' />
      <div>
        <Link href='/feed/configure/add'><FaPlus /></Link>
      </div>
      <ConfigureFeeds />
    </section>
  )
}