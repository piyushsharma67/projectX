import { Button, ButtonProps } from "@mui/material"
import React from "react"
import styles from './buttonStyle.module.css'

interface IButtonProps extends ButtonProps {
    buttontext: string
    buttontextstyle?: string
}

function ButtonComp(props: IButtonProps) {
    return (
        <Button
            {...props}
        >
            <span className={`${styles.buttonTextStyle} ${props.buttontextstyle}`}>{props.buttontext}</span>
        </Button>
    )
}

export default React.memo(ButtonComp)