import styles from "./Logo.module.css"

export function DOTsLogo() {
	return (
		<div className={styles.logoContainer}>
			<span className={styles.dot}>&bull;</span>
			<span className={styles.dotsText}>
				DOT<span style={{ textTransform: "lowercase" }}>s </span>
			</span>

			<span className={styles.dotsMeta}>boilerplates</span>
		</div>
	)
}
