import { PropsWithChildren } from "react"
import styles from "./styles.module.css"

export function Loading({ children, loading }: PropsWithChildren<{ loading: boolean }>) {
	if (!loading) {
		return <>{children}</>
	}

	return <div className={styles.load} />
}
