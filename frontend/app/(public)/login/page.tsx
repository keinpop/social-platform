import { LoginForm } from './form';
import Link from 'next/link';

export default function Page() {
  return (
    <div className="flex flex-col p-4 lg:w-1/3">
      <div className="text-center">
        <h1 className="text-3xl text-blue-950 font-bold">Вход</h1>
        <p className="text-gray-500">
          Введите почту и пароль, чтобы войти в аккаунт
        </p>
      </div>
      <div className="mt-6">
        <LoginForm />
      </div>
      <div className="mt-4 text-center text-blue-950 text-sm">
        Еще нет аккаунта?{' '}
        <Link className="underline text-blue-950" href="/signup">
          Регестрация
        </Link>
      </div>
    </div>
  );
}
