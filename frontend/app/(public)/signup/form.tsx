'use client';

import { Label } from '@/components/ui/label';
import { Input } from '@/components/ui/input';
import { Button } from '@/components/ui/button';
import { signup } from '@/app/auth/01-auth';
import { useFormState, useFormStatus } from 'react-dom';

export function SignupForm() {
  const [state, action] = useFormState(signup, undefined);

  return (
    <form action={action}>
      <div className="flex flex-col gap-2">
        <div>
          <Label htmlFor="name" className="text-blue-950">ФИО</Label>
          <Input id="name" name="name" placeholder="Иванов Иван Иванович" />
        </div>
        {state?.errors?.name && (
          <p className="text-sm text-red-500">{state.errors.name}</p>
        )}
        <div>
          <Label className="text-blue-950" htmlFor="email">Почта</Label>
          <Input id="email" name="email" placeholder="ivanovivan@example.com" />
        </div>
        {state?.errors?.email && (
          <p className="text-sm text-red-500">{state.errors.email}</p>
        )}
        <div>
          <Label className="text-blue-950" htmlFor="password">Пароль</Label>
          <Input id="password" name="password" type="password" />
        </div>
        {state?.errors?.password && (
          <div className="text-sm text-red-500">
            <p>Пароль должен:</p>
            <ul>
              {state.errors.password.map((error) => (
                <li key={error}>- {error}</li>
              ))}
            </ul>
          </div>
        )}
        <SignupButton />
      </div>
    </form>
  );
}

export function SignupButton() {
  const { pending } = useFormStatus();

  return (
    <Button aria-disabled={pending} type="submit" style={{backgroundColor: "#172554"}} className="mt-2 w-full">
      {pending ? 'Регистрация...' : 'Зарегистрироваться'}
    </Button>
  );
}
