import Link from 'next/link';
import { SignupForm } from '@/app/(public)/signup/form';
export default function Page() {
  return (
    <div className="flex flex-col p-4 lg:w-1/3">
      <div className="text-center">
        <h1 className="text-3xl font-bold text-blue-950">Создание аккаунта</h1>
        <p className="text-gray-500">Введите ваши данные для начала работы</p>
      </div>
      <div className="mt-6">
        <SignupForm />
      </div>
      <div className="mt-6 text-center text-sm">
        Уже есть аккаунт?{' '}
        <Link className="underline" href="/login">
          Вход
        </Link>
      </div>
    </div>
  );
}
