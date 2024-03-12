import React from 'react'
import { TextField } from '@mui/material'
import Column from '../../../../components/column/Column'
import styles from './styles/signUpForm.module.css'
import ButtonComp from '../../../../components/button/button'
import { FieldErrors, FieldValues, UseFormRegister } from 'react-hook-form'
import { AuthFormTypeObj } from '../../utils/formsEnums'

const sx = {
    margin: '20px 0'
}

interface SignUpFormProps {
    register: UseFormRegister<FieldValues>
    errors: FieldErrors<FieldValues>
    onClickSignUp: () => void
    onClickLogin: () => void
    selectedFormIndex: number
}

function AuthForm(props: SignUpFormProps) {

    return (
        <form style={{ padding: 0 }}>
            <Column style={styles.formContainer}>
                <p className={styles.header}>Lorem Ipsum is simply dummy text of the printing and typesetting industry.</p>
                {props.selectedFormIndex == AuthFormTypeObj.Signup ? <TextField
                    label="Enter Username"
                    variant="outlined"
                    className={styles.inputStyle}
                    sx={sx}
                    required
                    error={props.errors.name != null}
                    {...props.register('name', {
                        required: true
                    })}
                />
                    :
                    (null as unknown as React.ReactElement)
                }
                <TextField
                    label="Enter Email"
                    variant="outlined"
                    className={styles.inputStyle}
                    sx={sx}
                    type='email'
                    error={props.errors.email != null}
                    helperText={props.errors.email?.message?.toString()}
                    {...props.register('email', {
                        required: true,
                        pattern: /^([a-zA-Z0-9._%-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,})$/
                    })}
                />

                <TextField
                    label="Enter Password"
                    className={styles.inputStyle}
                    type="password"
                    sx={sx}
                    required
                    error={props.errors.password != null}
                    {...props.register('password', {
                        required: true
                    })}
                />
                {props.selectedFormIndex == AuthFormTypeObj.Signup ? <TextField
                    label="Confirm password"
                    className={styles.inputStyle}
                    type="password"
                    sx={sx}
                    required
                    error={props.errors.cnf_password != null}
                    {...props.register('cnf_password', {
                        required: true
                    })}
                />
                    :
                    (null as unknown as React.ReactElement)
                }
                {props.selectedFormIndex == AuthFormTypeObj.Signup ? <TextField
                    label="Enter Mobile"
                    className={styles.inputStyle}
                    sx={sx}
                    required
                    error={props.errors.mobile != null}
                    {...props.register('mobile', {
                        required: true
                    })}
                />
                    :
                    (null as unknown as React.ReactElement)
                }
                <div className={styles.buttonContainer}>
                    <ButtonComp
                        variant="contained"
                        className={styles.buttonStyle}
                        buttontext={props.selectedFormIndex == AuthFormTypeObj.Login ? "Login" : "Register"}
                        onClick={props.selectedFormIndex == AuthFormTypeObj.Login ? props.onClickLogin : props.onClickSignUp}
                    >
                    </ButtonComp>
                </div>
            </Column>
        </form>
    )
}

export default React.memo(AuthForm)