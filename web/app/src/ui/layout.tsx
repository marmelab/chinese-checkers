import React, {useEffect} from "react";
import {create} from "zustand";
import {Outlet} from "react-router-dom";
import {Menu} from "./menu";

/**
 * Layout data store.
 */
const useLayoutStore = create<{
	title: string;
}>(() => ({
	title: "",
}));

/**
 * Layout component.
 */
export function Layout({children}: React.PropsWithChildren<{}>) {
	const title = useLayoutStore(state => state.title);

	return (
		<>
			<header>
				<h1>{title}</h1>
			</header>

			{children ?? <Outlet/>}

			<Menu/>
		</>
	);
}

/**
 * Set the page title.
 */
export function usePageTitle(title: string): void {
	useEffect(() => {
		useLayoutStore.setState({title});
	}, [title]);
}
