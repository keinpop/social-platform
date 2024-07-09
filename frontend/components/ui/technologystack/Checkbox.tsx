import React from 'react';

interface Props {
  technology: string;
  isChecked: boolean;
  onChange: (technology: string) => void;
}

const Checkbox: React.FC<Props> = ({ technology, isChecked, onChange }) => {
  const handleCheckboxChange = () => {
    onChange(technology);
  };

  return (
    <label style={{ cursor: 'pointer' }}>
      <input type="checkbox" onChange={handleCheckboxChange} checked={isChecked} style={{marginLeft: '10px'}}/> {technology}
    </label>
  );
};

export default Checkbox;
