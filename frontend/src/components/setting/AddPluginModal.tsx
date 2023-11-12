import { css } from '@emotion/react'
import { Dialog, Button, Flex, Text, TextField, IconButton } from '@radix-ui/themes'
import { BiPlus } from 'react-icons/bi'

export const AddPluginModal = () => {
  const styles = {
    trigger: css({
      margin: '0 10px',
      cursor: 'pointer',
      fontSize: '20px',
    })
  }

  return (
    <Dialog.Root>
      <Dialog.Trigger>
        <IconButton radius='full' css={styles.trigger}><BiPlus /></IconButton>
      </Dialog.Trigger>

      <Dialog.Content>
        <Dialog.Title>Add Plugin</Dialog.Title>
        <Dialog.Description mb='4'></Dialog.Description>

        <Flex direction='column' gap='3'>
          <label>
            <Text as='div' size='2' mb='1' weight='bold'>Name</Text>
            <TextField.Input />
          </label>
          <label>
            <Text as='div' size='2' mb='1' weight='bold'>Command</Text>
            <TextField.Input />
          </label>
        </Flex>

        <Flex gap='3' mt='4' justify='end'>
          <Dialog.Close>
            <Button variant='soft' color='gray'>Cancel</Button>
          </Dialog.Close>
          <Dialog.Close>
            <Button>Save</Button>
          </Dialog.Close>
        </Flex>

      </Dialog.Content>
    </Dialog.Root>
  )
}