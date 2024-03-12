import style from './infoCard.module.css'

interface InfoCardProps {
    headingText: string,
    subHeadingText: string,
    children?: React.ReactElement
}

function InfoCard(props: InfoCardProps) {
    return (
        <div className={style.cardContainer}>
            <div className={style.iconContainer}>
                {props.children}
            </div>
            <span className={style.headingText}>{props.headingText}</span>
            <span className={style.subHeadingText}>{props.subHeadingText}</span>
        </div>
    )
}

export default InfoCard