import React from 'react'
import styles from './styles/studentImage.module.css'
import students from '../../../../assets/students.png'
import student2 from '../../../../assets/student2.png'
import { AuthFormTypeObj } from '../../utils/formsEnums'

interface StudentImageProps {
    index: number
}

function StudentImage(props: StudentImageProps) {
    return (
        <img src={props.index == AuthFormTypeObj.Login ? students : student2} alt="students" className={styles.image} />
    )
}
export default React.memo(StudentImage)