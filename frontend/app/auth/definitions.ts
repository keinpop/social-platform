import { z } from 'zod';

export const SignupFormSchema = z.object({
  name: z
    .string()
    .min(2, { message: 'Имя должно быть не менее 2 символов.' })
    .trim(),
  email: z.string().email({ message: 'Пожалуйста, введите действительную почту.' }).trim(),
  password: z
    .string()
    .min(8, { message: 'Быть не менее 8 символов в длину' })
    .regex(/[a-zA-Z]/, { message: 'Содержать хотя бы одну букву.' })
    .regex(/[0-9]/, { message: 'Содержать хотя бы одну цифру.' })
    .regex(/[^a-zA-Z0-9]/, {
      message: 'Содержать хотя бы один специальный символ.',
    })
    .trim(),
});

export const LoginFormSchema = z.object({
  email: z.string().email({ message: 'Пожалуйста, введите действительную почту.' }),
  password: z.string().min(1, { message: 'Пожалуйста, введите пароль.' }),
});

export type FormState =
  | {
      errors?: {
        name?: string[];
        email?: string[];
        password?: string[];
      };
      message?: string;
    }
  | undefined;

export type SessionPayload = {
  userId: string | number;
  expiresAt: Date;
};
