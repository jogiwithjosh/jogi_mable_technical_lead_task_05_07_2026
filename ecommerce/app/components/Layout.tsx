import { ReactNode } from "react";
import Header from "./Header";

interface Props {
    children: ReactNode;
}

export default function Layout({
    children
}: Props) {
    return (
        <>
            <Header />

            <main
                style={{
                    padding: "2rem"
                }}
            >
                {children}
            </main>
        </>
    );
}