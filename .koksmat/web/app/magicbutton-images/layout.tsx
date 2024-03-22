interface RootLayoutProps {
  children: React.ReactNode;
}

export default function RootLayout({ children }: RootLayoutProps) {
  return <div className="h-screen w-screen overflow-hidden">{children}</div>;
}
