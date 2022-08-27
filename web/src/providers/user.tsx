import { createContext, ReactNode, useContext, useEffect, useState } from "react"
import { FullUserFragment, useGetProfileQuery } from "./api"

interface UserContextShape {
	user: FullUserFragment | null
	loading: boolean
	setUser: (user: FullUserFragment | null) => void
}
export const UserContext = createContext<UserContextShape>({
	user: null,
	loading: false,
	setUser: () => {
		// noop
	},
})

export function useUser() {
	return useContext(UserContext)
}

export function UserProvider({ children }: { children: ReactNode }) {
	const [user, setUser] = useState<FullUserFragment | null>(null)
	const { data, loading } = useGetProfileQuery()

	useEffect(() => {
		if (data && data.getProfile) {
			setUser(data.getProfile)
		}
	}, [data])

	return <UserContext.Provider value={{ user, loading, setUser }}>{children}</UserContext.Provider>
}
