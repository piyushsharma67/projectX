import styles from './columnStyle.module.css'
import React from 'react'
interface ColumnProps {
    children?: React.ReactElement[]
    style?: React.CSSProperties | string
}

function Column(props: ColumnProps) {
    return (
        <div className={`${styles.columnContainer} ${props.style}`}>
            {props.children}
        </div>
    )
}

export default Column