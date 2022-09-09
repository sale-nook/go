import { PropsWithChildren, ReactNode } from "react";
import styles from "./DocsPage.module.css"

export function DocsPage({ children, sidebar }: PropsWithChildren<{sidebar: ReactNode}>) {
  return (
	<div className={styles.docsPage}>
	  <div className={styles.docsContent}>
		{children}
	  </div>
	  <aside className={styles.docsSidebar}>
	  	{sidebar}
		</aside>
	</div>
  )
}
