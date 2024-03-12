import React from 'react'
import FormToggle from "./formToggle/formToggle"
import AuthForm from "./authForm/authForm"
import StudentImage from "./sudentImage/studentImage"
import styles from './index.module.css'
import { FieldErrors, FieldValues, UseFormRegister } from "react-hook-form"
import { Navigate } from "react-router-dom"

interface ISignUpPageIndexProps {
    token: string
    loading: boolean
    forms: string[],
    errors: FieldErrors<FieldValues>
    register: UseFormRegister<FieldValues>,
    selectedFormIndex: number,
    onSelectedForm: (index: number) => () => void
    onClickLogin: () => void
    onClickSignUp: () => void
}

function SignUpPageIndex(props: ISignUpPageIndexProps) {
    return (
        <div className={styles.pageContainer}>
            <div className={styles.imageContainer}>
                <StudentImage index={props.selectedFormIndex} />
            </div>
            <div className={styles.formContainer}>
                <div className={styles.formToggleContainer}>
                    <FormToggle
                        forms={props.forms}
                        selectedFormIndex={props.selectedFormIndex}
                        onSelectedForm={props.onSelectedForm}
                    />
                </div>
                <div className={styles.form}>
                    <AuthForm
                        register={props.register}
                        errors={props.errors}
                        onClickSignUp={props.onClickSignUp}
                        onClickLogin={props.onClickLogin}
                        selectedFormIndex={props.selectedFormIndex}
                    />
                </div>
            </div>
            {props.token != "" ?? <Navigate to={'/Home'} replace={true} />}
        </div>
    )
}

export default SignUpPageIndex