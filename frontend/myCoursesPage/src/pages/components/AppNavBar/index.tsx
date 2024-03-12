import React from 'react'
import style from './appbar.module.css'
import logo from '../../../assets/logo.png'
import Button from '../../../components/button/button'
import { Avatar, useTheme } from '@mui/material'
import { useMemo } from 'react'
import { useNavigate } from 'react-router-dom'
import { useSelector } from 'react-redux'
import { RootAppState } from '../../../redux/store'

function AppNavBarIndex() {
    const theme = useTheme()
    const backgroundColor = theme.palette.secondary.main
    const registerButtonColor = theme.palette.info.main

    const { user } = useSelector((state: RootAppState) => state.authReducer)

    const buttonStyle = useMemo(() => ({
        backgroundColor: backgroundColor,
        borderRadius: 20
    }), [backgroundColor])

    const registerButtonStyle = useMemo(() => ({
        backgroundColor: registerButtonColor,
        borderRadius: 20
    }), [backgroundColor])

    const navigate = useNavigate()

    return (
        <div className={style.appBarContainer}>
            <img src={logo} alt="company logo" />
            <nav>
                <span className={style.navLinkButton}>Home</span>
                <span className={style.navLinkButton}>Courses</span>
                <span className={style.navLinkButton}>Careers</span>
                <span className={style.navLinkButton}>Blog</span>
                <span className={style.navLinkButton}>About us</span>
                {!user.name ? <>
                    <Button
                        buttontext='Login'
                        sx={registerButtonStyle}
                        buttontextstyle={style.registerButtonStyle}
                        onClick={() => {
                            navigate('/authenticate')
                        }}
                    />
                    <Button
                        buttontext='Register'
                        sx={buttonStyle}

                        onClick={() => {
                            navigate('/authenticate')
                        }}
                    /></> :
                    <Avatar data-testid="avatar" />
                }
            </nav>
        </div>
    )
}

export default AppNavBarIndex