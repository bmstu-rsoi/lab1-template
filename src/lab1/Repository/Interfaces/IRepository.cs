using System.Collections.Generic;
using System.Threading.Tasks;

namespace lab1.Repository.Interfaces
{
    public interface IRepository<T>
    {
        T Add(T element);
        List<T> GetAll();
        T Update(long Id, T element);
        int Delete(long id);
    }
}

