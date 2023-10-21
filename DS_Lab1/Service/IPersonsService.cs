using DS_Lab1.CommunicationElements;
using DS_Lab1.Models;

namespace DS_Lab1.Service
{
    public interface IPersonsService
    {
        IEnumerable<PersonResponse> GetAll();
        PersonResponse Get(int id);
        int Create(PersonRequest person);
        PersonResponse Update(int id, PersonRequest person);
        void Delete(int id);
    }
}
