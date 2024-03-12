import style from './block1.module.css'
import AppNavBarIndex from '../../../components/AppNavBar'
import { useTheme } from '@mui/material'
import image1 from '../../../../assets/img1.png'
import rectangle1 from '../../../../assets/rectangle1.png'
import rectangle2 from '../../../../assets/rectangle2.png'
import rectangle3 from '../../../../assets/rectangle3.png'
import rectangle4 from '../../../../assets/rectangle4.png'

function Landing1() {
    const theme = useTheme()
    const backgroundColor = theme.palette.primary.main
    return (
        <div className={style.block1Container} style={{ backgroundColor: backgroundColor }}>
            <AppNavBarIndex />
            <div className={style.landing1ContentStyle}>
                <div className={style.contentContainer}>
                    <p className={style.sloganText}><strong>Studying</strong> Online is now much easier</p>
                    <div className={style.text1Container}>
                        <p className={style.text1Style}>TOTC is an intresting platform that will teach you in more of an intresting way,</p>
                    </div>
                </div>
                <div className={style.allImagesContainer}>
                    <div className={style.imageContainer}>
                        <img src={image1} alt='image1' />
                    </div>
                    <div className={style.rectangle1Container}>
                        <img src={rectangle1} alt="rectangle1" />
                    </div>
                    <div className={style.rectangle2Container}>
                        <img src={rectangle2} alt="rectangle2" />
                    </div>
                    <div className={style.rectangle3Container}>
                        <img src={rectangle3} alt="rectangle2" />
                    </div>
                    <div className={style.rectangle4Container}>
                        <img src={rectangle4} alt="rectangle2" />
                    </div>
                </div>
            </div>
            {/* <div className={style.curve} /> */}
        </div>
    )
}

export default Landing1