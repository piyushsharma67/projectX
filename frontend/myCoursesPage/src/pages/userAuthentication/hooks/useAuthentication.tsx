import { useCallback, useState } from "react"
import { useForm } from "react-hook-form"
import { AuthFormTypeObj } from "../utils/formsEnums"
import { useAppDispatch } from "../../../redux/store"
import { LoginBody, SignUpBody, loginThunk, signUpThunk } from "../redux/thunks/authentication"
import { useNavigate } from "react-router-dom"

const forms = Object.keys(AuthFormTypeObj)

function useAuthentication() {
    const dispatch = useAppDispatch()
    const navigate = useNavigate()

    const [loading, setLoading] = useState(false)
    const [token, setToken] = useState('')
    const [selectedFormIndex, setSelectedFormIndex] = useState(AuthFormTypeObj.Login)

    const { register, handleSubmit, formState: { errors } } = useForm()

    const onSelectedForm = useCallback((index: number) => {
        return () => {
            setSelectedFormIndex(index)
        }
    }, [selectedFormIndex])

    const onClickSignUp = useCallback(() => {
        handleSubmit(async (data) => {
            try {
                const body: SignUpBody = {
                    email: data.email,
                    password: data.password,
                    name: data.name,
                    mobile: data.mobile
                }
                setLoading(true)
                const response = await dispatch(signUpThunk(body))
                setLoading(false)
                if (response.meta.requestStatus == 'rejected') {

                } else {
                    console.log("dsadasd")
                    navigate('/', { replace: true })
                }
            } catch (error) {

            }
        })()
    }, [])

    const onClickLogin = useCallback(() => {
        handleSubmit(async (data) => {
            console.log("handle", data)
            try {
                const body: LoginBody = {
                    email: data.email,
                    password: data.password,
                }
                setLoading(true)
                const response = await dispatch(loginThunk(body))
                setLoading(false)
                if (response.meta.requestStatus == 'rejected') {
                } else {
                    setToken("hi i am here")
                }
            } catch (error) {

            }
        })()
    }, [])

    return {
        token,
        loading,
        forms,
        errors,
        selectedFormIndex,
        register,
        onSelectedForm,
        onClickLogin,
        onClickSignUp
    }
}

export default useAuthentication