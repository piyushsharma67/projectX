import React from 'react'
import { Button } from '@mui/material'
import FormToggler from '../../../../components/formToggler/formToggler'
import styles from './formToggle.module.css'
import { AuthFormTypeObj } from '../../utils/formsEnums'

interface IFormToggleProps {
    forms: string[],
    selectedFormIndex: number,
    onSelectedForm: (index: number) => () => void
}

const buttonStyle = {
    borderRadius: 20,
    width: '100%',
    '&:focus': {
        outline: 'none', // Remove focus/active outline
    },
    '&:hover': {

    },
}

function FormToggle(props: IFormToggleProps) {

    return (
        <div className={styles.container}>
            {props.selectedFormIndex == AuthFormTypeObj.Login ? <p className={styles.heading}>Hi welcom to Loren Ipsum</p> : <p className={styles.heading}>sdas</p>}
            <div className={styles.toggleContainer}>
                <FormToggler >
                    <div className={styles.buttonContainer}>
                        {props.forms.map((text, index) => {
                            return (
                                <Button
                                    key={index}
                                    variant={props.selectedFormIndex - 1 == index ? "contained" : "text"}
                                    type="button"
                                    sx={buttonStyle}
                                    disableFocusRipple
                                    onClick={props.onSelectedForm(index + 1)}
                                    data-testid={`${text}-${index}`}
                                >
                                    <p className={styles.buttonText}>{text}</p>
                                </Button>
                            )
                        })}
                    </div>
                </FormToggler>
            </div>
        </div>
    )
}

export default React.memo(FormToggle)