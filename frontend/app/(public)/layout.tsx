import Link from 'next/link';
import { MenuIcon } from '@/components/ui/icons';

const links = [
  { href: 'https://www.mai.ru/', title: 'МАИ' },
  { href: 'https://t.me/LovePhysicsss', title: 'Здесь' },
  { href: 'https://t.me/thoughtscomefromthedarkness', title: 'Могли' },
  { href: 'https://t.me/Nyamerka', title: 'Быть' },
  { href: 'https://t.me/artempaskevichyan', title: 'Ваши' },
  { href: 'https://t.me/flaroteur', title: 'Кнопочки' },
];

export default function Layout({
  children,
}: Readonly<{
  children: React.ReactNode;
}>) {
  return (
    <div>
      <div className="border-b border-gray-100">
        <div className="container mx-auto flex max-w-7xl items-center justify-end p-4 md:justify-between md:px-6">
          <nav className="hidden items-center space-x-4 text-sm md:flex">
            {links.map((link) => (
              <Link className="text-gray-900" href={link.href} key={link.title}>
                {link.title}
              </Link>
            ))}
          </nav>
          <div className="hidden items-center space-x-4 md:flex">
            <Link
              className="inline-flex h-8 items-center text-blue-950 justify-center rounded-md border border-gray-200 bg-white px-4 text-sm font-medium shadow-sm transition-colors hover:bg-gray-100 hover:text-gray-900 focus-visible:outline-none focus-visible:ring-1 focus-visible:ring-gray-950 disabled:pointer-events-none disabled:opacity-5"
              href="/login"
            >
              Вход
            </Link>
          </div>
          <div className="flex items-center space-x-4 md:hidden">
            <Link
              className="inline-flex h-8 items-center rounded-md border border-gray-200 bg-white px-3 text-sm font-medium"
              href="/login"
            >
              Login
            </Link>
            <button className="inline-flex rounded-md md:hidden" type="button">
              <MenuIcon className="h-6 w-6" />
              <span className="sr-only">Toggle Menu</span>
            </button>
          </div>
        </div>
      </div>

      <main className="container mx-auto mt-36 flex max-w-7xl justify-center">
        {children}
      </main>
    </div>
  );
}
