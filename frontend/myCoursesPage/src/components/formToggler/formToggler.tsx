import React from 'react'
import { Box } from '@mui/material'
import styles from './formToggler.module.css'

interface FormTogglerProps {
    style?: React.CSSProperties | string,
    children?: React.ReactElement[] | React.ReactElement
}

const sx = {
    bgcolor: 'secondary.main'
}

function FormToggler(props: FormTogglerProps) {
    return (
        <Box
            className={`${styles.container} ${props.style}`}
            sx={sx}
        >
            {props.children}
        </Box>
    )
}

export default FormToggler