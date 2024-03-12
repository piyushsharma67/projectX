import styles from './block2.module.css'

function Block1() {
    return (
        <div className={styles.block2Container}>
            <div className={styles.headingContainer}>
                <p className={styles.heading}>Our Success</p>
                <p className={styles.subHeading}>Lorem Ipsum is simply dummy text of the printing and typesetting industry. Lorem Ipsum has been the industry's standard dummy text ever since the 1500s, when an unknown printer took a galley of type and scrambled it to make a type specimen book</p>

            </div>
            <div className={styles.numberContainer}>
                <div className={styles.textContainer}>
                    <span className={styles.gradientText}>15K+</span>
                    <span className={styles.subText}>students</span>
                </div>
                <div className={styles.textContainer}>
                    <span className={styles.gradientText}>75%</span>
                    <span className={styles.subText}>Total success</span>
                </div>
                <div className={styles.textContainer}>
                    <span className={styles.gradientText}>35</span>
                    <span className={styles.subText}>Main questions</span>
                </div>
                <div className={styles.textContainer}>
                    <span className={styles.gradientText}>26</span>
                    <span className={styles.subText}>Chief experts</span>
                </div>
                <div className={styles.textContainer}>
                    <span className={styles.gradientText}>16</span>
                    <span className={styles.subText}>Years of experience</span>
                </div>
            </div>
        </div>
    )
}

export default Block1