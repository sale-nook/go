import Link from "next/link"
import { PropsWithChildren } from "react"
import styles from "./DocBox.module.css"

interface DocBoxSectionProps {
	title: string
	description: string
}

interface DocBoxProps extends DocBoxSectionProps {
	href: string
	done?: boolean
}

export function DocBoxSection({ title, description, children }: PropsWithChildren<DocBoxSectionProps>) {
	return (
		<section className={styles.docBoxSection}>
			<h2 className={styles.title}>{title}</h2>
			<p className={styles.description}>{description}</p>
			<div className={styles.docBoxContainer}>{children}</div>
		</section>
	)
}

export function DocBox(props: DocBoxProps) {
	return (
		<div className={[styles.docBox, props.done ? styles.done : ""].join(" ")}>
			<Link href={props.href}>
				<a className={styles.docLink}>
					<h2 className={styles.docTitle}>{props.title}</h2>
					<p className={styles.docDescription}>{props.description}</p>
				</a>
			</Link>
		</div>
	)
}
