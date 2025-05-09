import React, {Suspense} from "react";
import {Outlet} from "react-router-dom";
import {Menu} from "./Menu";
import {LoaderView} from "./kit/Loader";

export function Layout({children}: React.PropsWithChildren<{}>) {
	return (
		<>
			<Suspense fallback={<LoaderView />}>{children ?? <Outlet />}</Suspense>

			<Menu />
		</>
	);
}
