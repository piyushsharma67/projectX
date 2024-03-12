
import styles from './block4.module.css'
import rectangle5 from '../../../../assets/rectangle5.png'

function Block4() {
    return (
        <div className={styles.block4Container}>
            <div className={styles.headingContainer}>
                <p className={styles.heading}>What is <span>TOTC</span></p>
                <span className={styles.subHeading}>TOTC is a platform that allows educators to create online classes whereby they can store the course materials online; manage assignments, quizzes and exams; monitor due dates; grade results and provide students with feedback all in one place.</span>
            </div>
            <div className={styles.textImageContainer}>
                <div className={styles.textContainer}>
                    <p className={styles.heading}>Everything you can do in a physical classroom, <span>you can do with TOTC</span></p>
                    <span className={styles.subHeading}>TOTC’s school management software helps traditional and online schools manage scheduling, attendance, payments and virtual classrooms all in one secure cloud-based system.</span>
                </div>
                <div className={styles.imageContainer}>
                    <img src={rectangle5} alt="rectangle5" />
                </div>

            </div>

        </div>
    )
}

export default Block4