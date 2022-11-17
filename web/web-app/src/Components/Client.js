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
            <a href='/res/app.apk' download>
                <Button variant='contained'>Download</Button>
            </a>
        </Box>
    )
}
