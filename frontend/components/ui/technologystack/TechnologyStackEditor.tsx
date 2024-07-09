import React, { useState, useEffect } from 'react';
import Checkbox from './Checkbox'; // Путь к вашему компоненту Checkbox

interface Props {
  selectedTechnologies: string[];
  onSave: (technologies: string[]) => void;
  validTechnologies: string[]; // Массив допустимых технологий
}

const TechnologyStackEditor: React.FC<Props> = ({ selectedTechnologies, onSave, validTechnologies }) => {
  const [showModal, setShowModal] = useState(false);
  const [localSelection, setLocalSelection] = useState<string[]>([]);
  const [isEditing, setIsEditing] = useState(false); // Состояние для отслеживания редактирования

  useEffect(() => {
    setLocalSelection(selectedTechnologies);
  }, [selectedTechnologies]);

  const toggleModal = () => {
    setShowModal(!showModal);
    setIsEditing(!showModal); // Переключаем состояние редактирования
  };

  const handleSave = () => {
    onSave(localSelection);
    toggleModal();
  };
  

  const handleCheckboxChange = (technology: string) => {
    const isSelected = localSelection ? localSelection.includes(technology) : false;
    if (isSelected) {
      setLocalSelection(localSelection.filter(tech => tech !== technology));
    } else {
      setLocalSelection([...(localSelection || []), technology]);
    }
  };

  return (
    <div>
      {!isEditing && <button onClick={toggleModal} style={{ fontSize: '14px', color: '#555'}}>Редактировать</button>}
      {showModal && (
        <div className="modal">
          <div className="modal-content">
            <h2>Выберите технологии</h2>
            {validTechnologies.map(tech => (
              <Checkbox
                key={tech}
                technology={tech}
                isChecked={localSelection.includes(tech)}
                onChange={handleCheckboxChange}
              />
            ))}

          </div>
          <button onClick={handleSave} style={{ fontSize: '14px', color: '#555'}}>Сохранить</button>
          <button onClick={toggleModal} style={{ marginLeft: '10px', fontSize: '14px', color: '#555' }}>Отмена</button>
        </div>
      )}
    </div>
  );
};

export default TechnologyStackEditor;
