import * as React from 'react'
import { Box, Button } from '@material-ui/core'

export default function Client () {
    return (
        <Box sx={{
            marginTop: 8,
            display: 'flex',
            flexDirection: 'column',
            alignItems: 'center'
          }}>
            {console.log('Current directory: ', __dirname)}
            <a href='clientWeb/res/app.apk' download>
            {/* <a href='test' download> */}
                <Button variant='contained'>Download</Button>
            </a>
        </Box>
    )
}
