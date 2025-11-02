import { useState } from 'react';
import { useNavigate } from 'react-router-dom';
import { Button, Input, Link, Modal, Select } from '../components';

export const Home = () => {
  const navigate = useNavigate();
  const [isModalOpen, setIsModalOpen] = useState(false);
  const [inputValue, setInputValue] = useState('');
  const [selectValue, setSelectValue] = useState('');

  return (
    <div className="min-h-screen bg-gray-900 text-white p-8">
      <div className="max-w-6xl mx-auto space-y-12">
        <header className="text-center">
          <h1 className="text-4xl font-bold text-indigo-400 mb-4">Компоненты Rentor</h1>
          <p className="text-gray-400">Демонстрация всех UI компонентов</p>
        </header>

        {/* Navigation Links */}
        <div className="flex gap-4 justify-center">
          <button
            onClick={() => navigate('/profile')}
            className="px-4 py-2 bg-gray-800 rounded-lg hover:bg-gray-700 transition-colors"
          >
            Profile
          </button>
          <button
            onClick={() => navigate('/advertisement/1')}
            className="px-4 py-2 bg-gray-800 rounded-lg hover:bg-gray-700 transition-colors"
          >
            Advertisement
          </button>
        </div>

        {/* Button Component */}
        <section className="bg-gray-800 rounded-lg p-6 space-y-6">
          <h2 className="text-2xl font-semibold text-indigo-400 mb-4">Button</h2>
          
          <div className="space-y-4">
            <div>
              <h3 className="text-lg text-gray-300 mb-3">Variants (Primary)</h3>
              <div className="flex flex-wrap gap-3">
                <Button variant="solid" colorScheme="primary">Solid</Button>
                <Button variant="outline" colorScheme="primary">Outline</Button>
                <Button variant="ghost" colorScheme="primary">Ghost</Button>
              </div>
            </div>

            <div>
              <h3 className="text-lg text-gray-300 mb-3">Variants (Neutral)</h3>
              <div className="flex flex-wrap gap-3">
                <Button variant="solid" colorScheme="neutral">Solid</Button>
                <Button variant="outline" colorScheme="neutral">Outline</Button>
                <Button variant="ghost" colorScheme="neutral">Ghost</Button>
              </div>
            </div>

            <div>
              <h3 className="text-lg text-gray-300 mb-3">Sizes</h3>
              <div className="flex flex-wrap items-center gap-3">
                <Button size="sm">Small</Button>
                <Button size="md">Medium</Button>
                <Button size="lg">Large</Button>
              </div>
            </div>

            <div>
              <h3 className="text-lg text-gray-300 mb-3">States</h3>
              <div className="flex flex-wrap gap-3">
                <Button disabled>Disabled</Button>
                <Button onClick={() => alert('Clicked!')}>Clickable</Button>
              </div>
            </div>
          </div>
        </section>

        {/* Link Component */}
        <section className="bg-gray-800 rounded-lg p-6 space-y-6">
          <h2 className="text-2xl font-semibold text-indigo-400 mb-4">Link</h2>
          
          <div className="space-y-4">
            <div>
              <h3 className="text-lg text-gray-300 mb-3">Variants (Primary)</h3>
              <div className="flex flex-wrap gap-4">
                <Link variant="default" colorScheme="primary" href="#">
                  Default Link
                </Link>
                <Link variant="underlined" colorScheme="primary" href="#">
                  Underlined Link
                </Link>
              </div>
            </div>

            <div>
              <h3 className="text-lg text-gray-300 mb-3">Variants (Neutral)</h3>
              <div className="flex flex-wrap gap-4">
                <Link variant="default" colorScheme="neutral" href="#">
                  Default Link
                </Link>
                <Link variant="underlined" colorScheme="neutral" href="#">
                  Underlined Link
                </Link>
              </div>
            </div>

            <div>
              <h3 className="text-lg text-gray-300 mb-3">Sizes</h3>
              <div className="flex flex-wrap items-center gap-4">
                <Link size="sm" href="#">Small</Link>
                <Link size="md" href="#">Medium</Link>
                <Link size="lg" href="#">Large</Link>
              </div>
            </div>

            <div>
              <h3 className="text-lg text-gray-300 mb-3">Router Link</h3>
              <Link to="/profile">Go to Profile</Link>
            </div>
          </div>
        </section>

        {/* Input Component */}
        <section className="bg-gray-800 rounded-lg p-6 space-y-6">
          <h2 className="text-2xl font-semibold text-indigo-400 mb-4">Input</h2>
          
          <div className="space-y-4">
            <div>
              <h3 className="text-lg text-gray-300 mb-3">Sizes (Primary)</h3>
              <div className="space-y-3 max-w-md">
                <Input size="sm" placeholder="Small input" colorScheme="primary" />
                <Input size="md" placeholder="Medium input" colorScheme="primary" />
                <Input size="lg" placeholder="Large input" colorScheme="primary" />
              </div>
            </div>

            <div>
              <h3 className="text-lg text-gray-300 mb-3">Controlled Input</h3>
              <div className="max-w-md">
                <Input
                  value={inputValue}
                  onChange={(e) => setInputValue(e.target.value)}
                  placeholder="Type something..."
                  colorScheme="primary"
                />
                <p className="text-sm text-gray-400 mt-2">Value: {inputValue || '(empty)'}</p>
              </div>
            </div>

            <div>
              <h3 className="text-lg text-gray-300 mb-3">Variants</h3>
              <div className="space-y-3 max-w-md">
                <Input placeholder="Default variant" variant="default" colorScheme="primary" />
                <Input placeholder="Error variant" variant="error" colorScheme="primary" />
              </div>
            </div>

            <div>
              <h3 className="text-lg text-gray-300 mb-3">States</h3>
              <div className="max-w-md">
                <Input placeholder="Disabled input" disabled className="mb-3" />
                <Input type="password" placeholder="Password input" colorScheme="primary" />
              </div>
            </div>
          </div>
        </section>

        {/* Select Component */}
        <section className="bg-gray-800 rounded-lg p-6 space-y-6">
          <h2 className="text-2xl font-semibold text-indigo-400 mb-4">Select</h2>
          
          <div className="space-y-4">
            <div>
              <h3 className="text-lg text-gray-300 mb-3">Sizes (Primary)</h3>
              <div className="space-y-3 max-w-md">
                <Select size="sm" colorScheme="primary">
                  <option value="">Small select</option>
                  <option value="1">Option 1</option>
                  <option value="2">Option 2</option>
                </Select>
                <Select size="md" colorScheme="primary">
                  <option value="">Medium select</option>
                  <option value="1">Option 1</option>
                  <option value="2">Option 2</option>
                </Select>
                <Select size="lg" colorScheme="primary">
                  <option value="">Large select</option>
                  <option value="1">Option 1</option>
                  <option value="2">Option 2</option>
                </Select>
              </div>
            </div>

            <div>
              <h3 className="text-lg text-gray-300 mb-3">Controlled Select</h3>
              <div className="max-w-md">
                <Select
                  value={selectValue}
                  onChange={(e) => setSelectValue(e.target.value)}
                  colorScheme="primary"
                >
                  <option value="">Choose an option</option>
                  <option value="option1">Option 1</option>
                  <option value="option2">Option 2</option>
                  <option value="option3">Option 3</option>
                </Select>
                <p className="text-sm text-gray-400 mt-2">Selected: {selectValue || '(none)'}</p>
              </div>
            </div>

            <div>
              <h3 className="text-lg text-gray-300 mb-3">Variants</h3>
              <div className="space-y-3 max-w-md">
                <Select variant="default" colorScheme="primary">
                  <option value="">Default variant</option>
                  <option value="1">Option 1</option>
                </Select>
                <Select variant="error" colorScheme="primary">
                  <option value="">Error variant</option>
                  <option value="1">Option 1</option>
                </Select>
              </div>
            </div>

            <div>
              <h3 className="text-lg text-gray-300 mb-3">States</h3>
              <div className="max-w-md">
                <Select disabled>
                  <option value="">Disabled select</option>
                </Select>
              </div>
            </div>
          </div>
        </section>

        {/* Modal Component */}
        <section className="bg-gray-800 rounded-lg p-6 space-y-6">
          <h2 className="text-2xl font-semibold text-indigo-400 mb-4">Modal</h2>
          
          <div className="space-y-4">
            <div>
              <Button onClick={() => setIsModalOpen(true)}>Open Modal</Button>
            </div>
            
            <p className="text-gray-400">
              Модальное окно закрывается по: клику вне модалки, нажатию ESC, клику на крестик
            </p>
          </div>

          <Modal isOpen={isModalOpen} onClose={() => setIsModalOpen(false)}>
            <div className="p-6">
              <h3 className="text-2xl font-semibold text-white mb-4">Модальное окно</h3>
              <p className="text-gray-300 mb-4">
                Это пример модального окна. Вы можете закрыть его несколькими способами:
              </p>
              <ul className="list-disc list-inside text-gray-300 space-y-2 mb-6">
                <li>Клик по крестику в правом верхнем углу</li>
                <li>Клик вне модального окна</li>
                <li>Нажатие клавиши ESC</li>
              </ul>
              <div className="flex gap-3">
                <Button onClick={() => setIsModalOpen(false)}>Закрыть</Button>
                <Button variant="outline" onClick={() => setIsModalOpen(false)}>
                  Отмена
                </Button>
              </div>
            </div>
          </Modal>
        </section>
      </div>
    </div>
  );
};

