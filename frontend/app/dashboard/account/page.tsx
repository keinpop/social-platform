'use client';
import { Button } from '@/components/ui/button';
import React , { useState } from 'react';
import {
  Card,
  CardContent,
  CardDescription,
  CardHeader,
  CardTitle,
} from '@/components/ui/card';
import EditableTextBlock from '@/components/ui/edibletextblock';
import { MainUserIcon } from '@/components/ui/icons';
import TechnologyStackEditor from '@/components/ui/technologystack/TechnologyStackEditor';


export default function Page() {

  const name = 'Имя пользователя'
  const humanDescription = 'Студент 2-го курса'

  const [currentTechnologies, setCurrentTechnologies] = useState<string[]>([]);
  const validTechnologies = [
    'JavaScript',
    'TypeScript',
    'React',
    'Vue.js',
    'Next.js',
    'Nuxt.js',
    'Python',
    'Django',
    'Flask',
    'Golang',
    'Java',
    'C++',
    'C#',
    'PHP',
    'Ruby',
  ]

  const handleSaveTechnologies = (technologies: string[]) => {
    setCurrentTechnologies(technologies);
    // Здесь вы можете выполнить дополнительные действия по сохранению, например, отправить на сервер
    console.log('Сохранены на главной странице:', technologies);
  };

  return (
    <div className="grid gap-4 md:gap-8">
      <Card >
        <CardHeader>
          <CardTitle>Личный кабинет</CardTitle>
        </CardHeader>
        <CardContent>
          <div className="flex justify-between">
            <Card className='' style={{ marginRight: '24px', maxWidth: 'calc(50% - 10px)' }}>
              <CardHeader>
                  <MainUserIcon/>
              </CardHeader>
            </Card>
            <Card className='' style={{ flex: 1 }}>
              <CardHeader>
                <CardTitle>{name}</CardTitle>
                <CardDescription style={{ marginBottom: '25px' }}>{humanDescription}</CardDescription>
                <EditableTextBlock defaultText={'Описание пользователя'}/>
              </CardHeader>
            </Card>
          </div>
          <Card className='mt-5'>
            <p>Cтек технологий: {currentTechnologies.join(', ')}</p>
            <TechnologyStackEditor 
              selectedTechnologies={currentTechnologies} 
              onSave={handleSaveTechnologies} 
              validTechnologies={validTechnologies} />
          </Card>
        </CardContent>
      </Card>
    </div>
  );
}
